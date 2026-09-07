import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ClinicalOrderService } from '../../../services/ClinicalOrder.service';
import { ClinicalOrder } from '../../../models/ClinicalOrder';
import { SubBaseComponent } from '../../ClinicalOrder/sub.base.component';

@Component({
    selector: 'app-create-clinicalOrder',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateClinicalOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Add ClinicalOrder';

    clinicalOrderForm: FormGroup;
    clinicalOrder: ClinicalOrder;

    constructor( http: HttpClient,
        private clinicalOrderService: ClinicalOrderService,
        private fb: FormBuilder,
        private router: Router
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

    
    addClinicalOrder(orderNumber, Patient, Encounter, OrderingClinician, MedicationOrders, LaboratoryOrders, ImagingOrders, ProcedureOrders, Authorizations, Status, OrderType, Priority): void {
        this.clinicalOrderService
        .addClinicalOrder(orderNumber, Patient, Encounter, OrderingClinician, MedicationOrders, LaboratoryOrders, ImagingOrders, ProcedureOrders, Authorizations, Status, OrderType, Priority)
            .subscribe(() => {
                this.router.navigate(['/indexClinicalOrder']);
            });
    }

    ngOnInit(): void {
    }
}