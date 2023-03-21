package com.ecommerce.orderservice.repository;

import com.ecommerce.orderservice.model.Order;
import org.springframework.data.jpa.repository.JpaRepository;

public interface OrderReposity extends JpaRepository<Order, Long> {
}
