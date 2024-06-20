import express, { Request, Response } from 'express';
import { exec } from 'child_process';
import { promisify } from 'util';

const app = express();
const port = process.env.PORT || 3000;

const execAsync = promisify(exec);

app.use(express.json());

app.post('/run-command', async (req: Request, res: Response) => {
  const { command } = req.body;

  if (!command) {
    return res.status(400).send({ error: 'Command is required' });
  }

  try {
    const { stdout, stderr } = await execAsync(command);
    console.log("command:::::::::::::::::::: ", command);
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
