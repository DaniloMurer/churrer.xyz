package xyz.churrer.house.service;

import jakarta.enterprise.context.ApplicationScoped;
import jakarta.persistence.EntityManager;
import jakarta.transaction.Transactional;
import xyz.churrer.house.domain.jpa.Telemetry;

import java.util.List;

/**
 * Service class for database interactions with telemetries
 */
@ApplicationScoped
public class TelemetryService {

    private final EntityManager em;

    public TelemetryService(EntityManager em) {
        this.em = em;
    }

    /**
     * Save telemetry to database
     * @param telemetry {@link Telemetry} to persist
     */
    @Transactional
    public void persist(Telemetry telemetry) {
        em.persist(telemetry);
    }

    /**
     * Retrieve all telemetries from database
     * @return {@link List} of {@link Telemetry}
     */
    public List<Telemetry> findAll() {
        return em.createQuery("select t from Telemetry t", Telemetry.class).getResultList();
    }
}
