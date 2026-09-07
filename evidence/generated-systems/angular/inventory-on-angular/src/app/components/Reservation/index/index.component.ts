
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ReservationService } from '../../../services/Reservation.service';
import { Reservation } from '../../../models/Reservation';

@Component({
    selector: 'app-index-reservation',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexReservationComponent implements OnInit {

    reservations: Reservation[] = [];

    constructor(
        private router: Router,
        private service: ReservationService
) {}

    ngOnInit(): void {
        this.getReservations();
}

    getReservations(): void {
        this.service.getReservations().subscribe((res) => {
        this.reservations = res;
    });
}

    deleteReservation(id: any): void {
        this.service.deleteReservation(id)
            .subscribe(() => {
                this.getReservations();
            });
    }
}