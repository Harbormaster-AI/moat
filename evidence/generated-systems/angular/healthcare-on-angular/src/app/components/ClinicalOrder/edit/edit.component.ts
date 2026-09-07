import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ClinicalOrderService } from '../../../services/ClinicalOrder.service';
import { SubBaseComponent } from '../../ClinicalOrder/sub.base.component';


@Component({
    selector: 'app-edit-clinicalOrder',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditClinicalOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Edit ClinicalOrder';

    clinicalOrderForm: FormGroup;
    clinicalOrder: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ClinicalOrderService,
        private fb: FormBuilder
) {
        super(http);
        this.clinicalOrderForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  orderNumber: ['', Validators.required],
      Patient: ['', ],
      Encounter: ['', ],
      OrderingClinician: ['', ],
      MedicationOrders: ['', ],
      LaboratoryOrders: ['', ],
      ImagingOrders: ['', ],
      ProcedureOrders: ['', ],
      Authorizations: ['', ],
      Status: ['', ],
      OrderType: ['', ],
      Priority: ['', ]
        });
    }

    
    updateClinicalOrder(orderNumber, Patient, Encounter, OrderingClinician, MedicationOrders, LaboratoryOrders, ImagingOrders, ProcedureOrders, Authorizations, Status, OrderType, Priority): void {
        this.route.params.subscribe((params) => {

                        this.service.updateClinicalOrder(orderNumber, Patient, Encounter, OrderingClinician, MedicationOrders, LaboratoryOrders, ImagingOrders, ProcedureOrders, Authorizations, Status, OrderType, Priority, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexClinicalOrder']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getClinicalOrder(params['id']).subscribe(res => {
                this.clinicalOrder = res;
            });
        });
    }
}