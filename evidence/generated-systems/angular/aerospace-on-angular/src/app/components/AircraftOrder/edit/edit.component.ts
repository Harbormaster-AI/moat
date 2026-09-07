import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AircraftOrderService } from '../../../services/AircraftOrder.service';
import { SubBaseComponent } from '../../AircraftOrder/sub.base.component';


@Component({
    selector: 'app-edit-aircraftOrder',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAircraftOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Edit AircraftOrder';

    aircraftOrderForm: FormGroup;
    aircraftOrder: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: AircraftOrderService,
        private fb: FormBuilder
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

    
    updateAircraftOrder(orderNumber, totalAmount, Operator, Variant, Quote, PurchaseAgreement, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAircraftOrder(orderNumber, totalAmount, Operator, Variant, Quote, PurchaseAgreement, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAircraftOrder']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAircraftOrder(params['id']).subscribe(res => {
                this.aircraftOrder = res;
            });
        });
    }
}