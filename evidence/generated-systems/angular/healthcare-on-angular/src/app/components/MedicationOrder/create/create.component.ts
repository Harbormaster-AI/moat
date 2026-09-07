import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MedicationOrderService } from '../../../services/MedicationOrder.service';
import { MedicationOrder } from '../../../models/MedicationOrder';
import { SubBaseComponent } from '../../MedicationOrder/sub.base.component';

@Component({
    selector: 'app-create-medicationOrder',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateMedicationOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Add MedicationOrder';

    medicationOrderForm: FormGroup;
    medicationOrder: MedicationOrder;

    constructor( http: HttpClient,
        private medicationOrderService: MedicationOrderService,
        private fb: FormBuilder,
        private router: Router
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

    
    addMedicationOrder(medicationCode, dose, frequency, duration, Order, Pharmacy, Dispenses, Route): void {
        this.medicationOrderService
        .addMedicationOrder(medicationCode, dose, frequency, duration, Order, Pharmacy, Dispenses, Route)
            .subscribe(() => {
                this.router.navigate(['/indexMedicationOrder']);
            });
    }

    ngOnInit(): void {
    }
}