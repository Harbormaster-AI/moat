import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ReservationService } from '../../../services/Reservation.service';
import { Reservation } from '../../../models/Reservation';
import { SubBaseComponent } from '../../Reservation/sub.base.component';

@Component({
    selector: 'app-create-reservation',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateReservationComponent extends SubBaseComponent implements OnInit {

    title = 'Add Reservation';

    reservationForm: FormGroup;
    reservation: Reservation;

    constructor( http: HttpClient,
        private reservationService: ReservationService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.reservationForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  referenceNumber: ['', Validators.required],
      reservedQuantity: ['', Validators.required],
      promisedDate: ['', Validators.required],
      Sku: ['', ],
      Warehouse: ['', ],
      Location: ['', ],
      InventoryItem: ['', ],
      Lot: ['', ],
      SerialNumbers: ['', ],
      DemandSignal: ['', ],
      ReservationStatus: ['', ],
      ReservationType: ['', ]
        });
    }

    
    addReservation(referenceNumber, reservedQuantity, promisedDate, Sku, Warehouse, Location, InventoryItem, Lot, SerialNumbers, DemandSignal, ReservationStatus, ReservationType): void {
        this.reservationService
        .addReservation(referenceNumber, reservedQuantity, promisedDate, Sku, Warehouse, Location, InventoryItem, Lot, SerialNumbers, DemandSignal, ReservationStatus, ReservationType)
            .subscribe(() => {
                this.router.navigate(['/indexReservation']);
            });
    }

    ngOnInit(): void {
    }
}