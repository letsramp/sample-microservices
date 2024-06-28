import express, { Request, Response } from 'express';
import { exec } from 'child_process';
import { promisify } from 'util';

const app = express();
const port = process.env.PORT || 4000;

const execAsync = promisify(exec);

app.use(express.json());

app.post('/run-command', async (req: Request, res: Response) => {
  const { command } = req.body;

  console.log("command:::::::::::::::::::: ", command);

  if (!command) {
    return res.status(400).send({ error: 'Command is required' });
  }

  let cmds = "mkdir -p ../skyramp-agent && cd ../skyramp-agent && output=\$("+command+") && filename=$(echo \"$output\" | grep -o 'mocks/mock-[^ ]*.yaml' | awk -F'/' '{print $NF}') && skyramp mocker apply \$filename --address localhost:35142 && gh codespace ports visibility 60000:public -c $CODESPACE_NAME"

  try {
    const { stdout, stderr } = await execAsync(cmds);
    console.log("cmds start:::::::::::::::::::: ", cmds);
    console.log(stdout)
    console.log(stderr)
    console.log("cmds end:::::::::::::::::::: ");
    res.status(200).send({
      success: true,
      output: stdout,
      error: stderr
    });
  } catch (error) {
    res.status(500).send({
      success: false,
      error: (error as Error).message
    });
  }
});

app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});
