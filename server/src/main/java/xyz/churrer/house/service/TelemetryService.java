package xyz.churrer.house.service;

import jakarta.enterprise.context.ApplicationScoped;
import jakarta.inject.Inject;
import jakarta.persistence.EntityManager;
import jakarta.transaction.Transactional;
import xyz.churrer.house.domain.jpa.Telemetry;

@ApplicationScoped
public class TelemetryService {
    @Inject
    EntityManager em;

    @Transactional
    public void persist(Telemetry telemetry) {
        em.persist(telemetry);
    }
}
