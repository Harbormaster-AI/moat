
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ObservationService } from '../../../services/Observation.service';
import { Observation } from '../../../models/Observation';

@Component({
    selector: 'app-index-observation',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexObservationComponent implements OnInit {

    observations: Observation[] = [];

    constructor(
        private router: Router,
        private service: ObservationService
) {}

    ngOnInit(): void {
        this.getObservations();
}

    getObservations(): void {
        this.service.getObservations().subscribe((res) => {
        this.observations = res;
    });
}

    deleteObservation(id: any): void {
        this.service.deleteObservation(id)
            .subscribe(() => {
                this.getObservations();
            });
    }
}