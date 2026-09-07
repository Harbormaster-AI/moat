
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { FlightHealthEventService } from '../../../services/FlightHealthEvent.service';
import { FlightHealthEvent } from '../../../models/FlightHealthEvent';

@Component({
    selector: 'app-index-flightHealthEvent',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexFlightHealthEventComponent implements OnInit {

    flightHealthEvents: FlightHealthEvent[] = [];

    constructor(
        private router: Router,
        private service: FlightHealthEventService
) {}

    ngOnInit(): void {
        this.getFlightHealthEvents();
}

    getFlightHealthEvents(): void {
        this.service.getFlightHealthEvents().subscribe((res) => {
        this.flightHealthEvents = res;
    });
}

    deleteFlightHealthEvent(id: any): void {
        this.service.deleteFlightHealthEvent(id)
            .subscribe(() => {
                this.getFlightHealthEvents();
            });
    }
}