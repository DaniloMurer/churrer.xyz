package xyz.churrer.house.domain.jpa;

import jakarta.persistence.*;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.time.LocalDateTime;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Entity
public class Telemetry {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    @Column(nullable = false)
    private String countryName;
    @Column(nullable = false)
    private String countryIso;
    @Column(nullable = false)
    private LocalDateTime timestamp;

    public Telemetry(String countryName, String countryIso, LocalDateTime timestamp) {
        this.countryName = countryName;
        this.countryIso = countryIso;
        this.timestamp = timestamp;
    }
}
