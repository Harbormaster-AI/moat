import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AircraftOrderService } from '../../../services/AircraftOrder.service';
import { AircraftOrder } from '../../../models/AircraftOrder';
import { SubBaseComponent } from '../../AircraftOrder/sub.base.component';

@Component({
    selector: 'app-create-aircraftOrder',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateAircraftOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Add AircraftOrder';

    aircraftOrderForm: FormGroup;
    aircraftOrder: AircraftOrder;

    constructor( http: HttpClient,
        private aircraftOrderService: AircraftOrderService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.aircraftOrderForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  orderNumber: ['', Validators.required],
      totalAmount: ['', Validators.required],
      Operator: ['', ],
      Variant: ['', ],
      Quote: ['', ],
      PurchaseAgreement: ['', ],
      Status: ['', ]
        });
    }

    
    addAircraftOrder(orderNumber, totalAmount, Operator, Variant, Quote, PurchaseAgreement, Status): void {
        this.aircraftOrderService
        .addAircraftOrder(orderNumber, totalAmount, Operator, Variant, Quote, PurchaseAgreement, Status)
            .subscribe(() => {
                this.router.navigate(['/indexAircraftOrder']);
            });
    }

    ngOnInit(): void {
    }
}