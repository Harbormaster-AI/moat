import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { MedicationOrderService } from '../../../services/MedicationOrder.service';
import { SubBaseComponent } from '../../MedicationOrder/sub.base.component';


@Component({
    selector: 'app-edit-medicationOrder',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditMedicationOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Edit MedicationOrder';

    medicationOrderForm: FormGroup;
    medicationOrder: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: MedicationOrderService,
        private fb: FormBuilder
) {
        super(http);
        this.medicationOrderForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  medicationCode: ['', Validators.required],
      dose: ['', Validators.required],
      frequency: ['', Validators.required],
      duration: ['', Validators.required],
      Order: ['', ],
      Pharmacy: ['', ],
      Dispenses: ['', ],
      Route: ['', ]
        });
    }

    
    updateMedicationOrder(medicationCode, dose, frequency, duration, Order, Pharmacy, Dispenses, Route): void {
        this.route.params.subscribe((params) => {

                        this.service.updateMedicationOrder(medicationCode, dose, frequency, duration, Order, Pharmacy, Dispenses, Route, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexMedicationOrder']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getMedicationOrder(params['id']).subscribe(res => {
                this.medicationOrder = res;
            });
        });
    }
}