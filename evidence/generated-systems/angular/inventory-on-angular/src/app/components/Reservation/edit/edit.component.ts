import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ReservationService } from '../../../services/Reservation.service';
import { SubBaseComponent } from '../../Reservation/sub.base.component';


@Component({
    selector: 'app-edit-reservation',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditReservationComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Reservation';

    reservationForm: FormGroup;
    reservation: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ReservationService,
        private fb: FormBuilder
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

    
    updateReservation(referenceNumber, reservedQuantity, promisedDate, Sku, Warehouse, Location, InventoryItem, Lot, SerialNumbers, DemandSignal, ReservationStatus, ReservationType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateReservation(referenceNumber, reservedQuantity, promisedDate, Sku, Warehouse, Location, InventoryItem, Lot, SerialNumbers, DemandSignal, ReservationStatus, ReservationType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexReservation']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getReservation(params['id']).subscribe(res => {
                this.reservation = res;
            });
        });
    }
}