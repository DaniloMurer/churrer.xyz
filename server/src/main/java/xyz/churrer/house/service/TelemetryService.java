package xyz.churrer.house.service;

import jakarta.enterprise.context.ApplicationScoped;
import jakarta.inject.Inject;
import jakarta.persistence.EntityManager;
import jakarta.transaction.Transactional;
import xyz.churrer.house.domain.jpa.Telemetry;

import java.util.List;

@ApplicationScoped
public class TelemetryService {
    @Inject
    EntityManager em;

    @Transactional
    public void persist(Telemetry telemetry) {
        em.persist(telemetry);
    }

    public List<Telemetry> findAll() {
        return em.createQuery("select t from Telemetry t", Telemetry.class).getResultList();
    }
}
