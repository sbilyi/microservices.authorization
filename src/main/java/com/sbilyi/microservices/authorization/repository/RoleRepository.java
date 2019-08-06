package com.sbilyi.microservices.authorization.repository;

import com.sbilyi.microservices.authorization.model.Role;
import com.sbilyi.microservices.authorization.model.RoleName;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.Optional;

@Repository
public interface RoleRepository extends JpaRepository<Role, Long> {
    Optional<Role> findByName(RoleName roleName);
}