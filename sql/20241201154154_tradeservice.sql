-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `trade_order` (
    id BIGINT PRIMARY KEY COMMENT '交易单id',
    biz_type INT NOT NULL COMMENT '业务类型，如DouTok充值、商城订单等',
    sale_way INT NOT NULL COMMENT '销售方式，如直接销售、代理销售等',
    pay_sequence INT NOT NULL COMMENT '支付顺序, 如先支付、后支付',
    goods_type INT NOT NULL COMMENT '商品类型，如普通商品、虚拟商品、优惠券等',
    status INT NOT NULL COMMENT '交易单状态，如已创建、已支付、开始交付、交付完成等',
    biz_status BIGINT NOT NULL DEFAULT 0 COMMENT '业务状态',
    biz_order_no varchar(64) DEFAULT '' NOT NULL COMMENT '业务订单号',
    biz_id BIGINT NOT NULL DEFAULT 0 COMMENT '业务id',
    buyer_id BIGINT NOT NULL COMMENT '买家id',
    seller_id BIGINT NOT NULL COMMENT '卖家id',
    pay_type INT NOT NULL COMMENT '支付类型，如余额支付、微信支付、支付宝支付等',
    payable_amount BIGINT NOT NULL COMMENT '应付金额，单位：厘',
    paid_amount BIGINT NOT NULL COMMENT '已支付金额，单位：厘',
    promotion_amount BIGINT NOT NULL COMMENT '优惠金额，单位：厘',
    delivery_info_id BIGINT NOT NULL DEFAULT 0 COMMENT '发货信息id',
    delivery_time TIMESTAMP COMMENT '发货时间',
    payment_info_id BIGINT NOT NULL DEFAULT 0 COMMENT '支付信息id',
    pay_time TIMESTAMP COMMENT '支付时间',
    close_time TIMESTAMP COMMENT '关单时间',
    close_type INT NOT NULL DEFAULT 0 COMMENT '关单类型',
    close_reason TEXT NOT NULL COMMENT '关单原因',
    extra JSON NOT NULL COMMENT '额外信息',
    trade_order_relation INT NOT NULL DEFAULT 0 COMMENT '交易单类型',
    origin_order_id BIGINT NOT NULL DEFAULT 0 COMMENT '原始订单id',
    expect_pay_time_out TIMESTAMP COMMENT '预期支付超时时间',
    actual_pay_time_out TIMESTAMP COMMENT '实际支付超时时间',
    execute_after_pay_time_out BOOLEAN NOT NULL DEFAULT FALSE COMMENT '支付超时后是否执行',
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `delivery_info` (
    id BIGINT PRIMARY KEY COMMENT '发货信息id',
    account_id BIGINT NOT NULL COMMENT '所属账户id',
    receiver_name varchar(64) DEFAULT '' NOT NULL COMMENT '收货人姓名',
    receiver_phone varchar(64) DEFAULT '' NOT NULL COMMENT '收货人电话',
    receiver_address TEXT NOT NULL COMMENT '收货地址',
    extra JSON NOT NULL COMMENT '额外信息',
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `payment_info` (
    id BIGINT PRIMARY KEY COMMENT '支付信息id',
    trade_order_id BIGINT NOT NULL COMMENT '交易单id',
    payment_no varchar(64) DEFAULT '' NOT NULL COMMENT '支付单号',
    pay_scene INT NOT NULL COMMENT '支付场景，如下单支付、退款支付等',
    paid_amount BIGINT NOT NULL COMMENT '支付金额，单位：厘',
    pay_channel INT NOT NULL DEFAULT 0 COMMENT '支付渠道，如余额支付、微信支付、支付宝支付等',
    pay_status INT NOT NULL COMMENT '支付状态，如未支付、已支付、支付失败等',
    pay_time TIMESTAMP COMMENT '支付时间',
    close_time TIMESTAMP COMMENT '关单时间',
    success_time TIMESTAMP COMMENT '支付成功时间',
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `refund_order` (
    id BIGINT PRIMARY KEY COMMENT '退款单id',
    trade_order_id BIGINT NOT NULL COMMENT '交易单id',
    payment_id BIGINT NOT NULL COMMENT '支付信息id',
    refund_no varchar(64) DEFAULT '' NOT NULL COMMENT '退款单号',
    refund_amount BIGINT NOT NULL COMMENT '退款金额，单位：厘',
    refund_status INT NOT NULL COMMENT '退款状态，如未退款、已退款、退款失败等',
    refund_type INT NOT NULL COMMENT '退款类型，如仅退款、退货退款等',
    refund_reason TEXT NOT NULL COMMENT '退款原因',
    refund_time TIMESTAMP COMMENT '退款时间',
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    close_time TIMESTAMP COMMENT '关单时间',
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `promotion` (
    id BIGINT PRIMARY KEY COMMENT '优惠id',
    name varchar(64) DEFAULT '' NOT NULL COMMENT '优惠名称',
    description TEXT NOT NULL COMMENT '优惠描述',
    promotion_type INT NOT NULL COMMENT '优惠类型，如满减、折扣等',
    promotion_rule JSON NOT NULL COMMENT '优惠规则',
    promotion_status INT NOT NULL COMMENT '优惠状态，如未开始、进行中、已结束等',
    start_time TIMESTAMP COMMENT '优惠开始时间',
    end_time TIMESTAMP COMMENT '优惠结束时间',

    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `trade_promotion` (
    id BIGINT PRIMARY KEY COMMENT '具体优惠id',
    promotion_id BIGINT NOT NULL COMMENT '优惠id',
    amount BIGINT NOT NULL COMMENT '优惠金额，单位：厘',
    promotion_type INT NOT NULL COMMENT '优惠类型，如满减、折扣等',
    account_id BIGINT NOT NULL COMMENT '账户id',

    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `trade_order`;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS `delivery_info`;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS `payment_info`;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS `refund_order`;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS `promotion`;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS `trade_promotion`;
-- +goose StatementEnd