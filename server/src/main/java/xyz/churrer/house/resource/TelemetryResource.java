package xyz.churrer.house.resource;

import jakarta.ws.rs.*;
import jakarta.ws.rs.core.MediaType;
import xyz.churrer.house.domain.dto.TelemetryDto;
import xyz.churrer.house.service.TelemetryService;
import java.util.List;

@Path("/api/telemetry")
public class TelemetryResource {

    private final TelemetryService telemetryService;

    public TelemetryResource(TelemetryService telemetryService) {
        this.telemetryService = telemetryService;
    }

    @POST
    public String create(TelemetryDto telemetry) {
        this.telemetryService.persist(telemetry.toTelemetry());
        return telemetry.toString();
    }

    @GET
    @Produces(MediaType.APPLICATION_JSON)
    public List<TelemetryDto> getAll() {
        return this.telemetryService.findAll().stream().map(TelemetryDto::fromTelemetry).toList();
    }
}
