package xyz.churrer.house.resource;

import jakarta.inject.Inject;
import jakarta.ws.rs.*;
import jakarta.ws.rs.core.MediaType;
import xyz.churrer.house.domain.dto.TelemetryDto;
import xyz.churrer.house.service.TelemetryService;

import java.util.List;
import java.util.stream.Collectors;

@Path("/api/telemetry")
public class TelemetryResource {

    @Inject
    TelemetryService telemetryService;

    @POST
    public String create(TelemetryDto telemetry) {
        this.telemetryService.persist(telemetry.toTelemetry());
        return telemetry.toString();
    }

    @GET
    @Produces(MediaType.APPLICATION_JSON)
    public List<TelemetryDto> getAll() {
        return this.telemetryService.findAll().stream().map(TelemetryDto::fromTelemetry).collect(Collectors.toList());
    }
}
