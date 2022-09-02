package main

import (
	pb "cartservice/genproto"
)

var db = make(map[string][]*pb.CartItem)

func AddItem(user string, productId string, quantity int32) error {
	items, found := db[user]
	if !found {
		items = make([]*pb.CartItem, 0)
	}
	var existingItem *pb.CartItem
	for _, cartItem := range items {
		if productId == cartItem.ProductId {
			existingItem = cartItem
		}
	}
	if existingItem == nil {
		existingItem = &pb.CartItem{ProductId: productId, Quantity: quantity}
		items = append(items, existingItem)
	} else {
		existingItem.Quantity += quantity
	}
	db[user] = items
	return nil
}

func EmtyCart(userId string) {
	db[userId] = make([]*pb.CartItem, 0)
}

func GetCart(userId string) *pb.Cart {
	cart := &pb.Cart{UserId: userId}
	items, found := db[userId]
	if found {
		cart.Items = items
	} else {
		cart.Items = make([]*pb.CartItem, 0)
	}
	return cart
}
