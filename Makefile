.PHONY: gen prepare move wire

OUT_DIR := ./output
ORDER_DIR := ./order
API_DIR := ./api

move:
	rm -rf $(API_DIR)/kitex_gen
	mv kitex_gen $(API_DIR)/

	rm -rf $(ORDER_DIR)/main.go $(ORDER_DIR)/kitex_info.yaml $(ORDER_DIR)/handler.go $(ORDER_DIR)/build.sh $(ORDER_DIR)/script

	mv main.go kitex_info.yaml handler.go build.sh script $(ORDER_DIR)/
	@echo "done"

kitex_gen:
	kitex -service order -module eshop -type protobuf -I api/order/ ./api/order/order.proto

gen: prepare api/kitex_gen move

prepare:
	mkdir -p $(OUT_DIR)

wire:
	@cd bff
	wire
	@cd ..

