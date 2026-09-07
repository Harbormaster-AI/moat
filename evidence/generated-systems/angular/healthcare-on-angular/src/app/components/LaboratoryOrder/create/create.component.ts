import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { LaboratoryOrderService } from '../../../services/LaboratoryOrder.service';
import { LaboratoryOrder } from '../../../models/LaboratoryOrder';
import { SubBaseComponent } from '../../LaboratoryOrder/sub.base.component';

@Component({
    selector: 'app-create-laboratoryOrder',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateLaboratoryOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Add LaboratoryOrder';

    laboratoryOrderForm: FormGroup;
    laboratoryOrder: LaboratoryOrder;

    constructor( http: HttpClient,
        private laboratoryOrderService: LaboratoryOrderService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.laboratoryOrderForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  testCode: ['', Validators.required],
      fastingRequired: ['', Validators.required],
      Order: ['', ],
      Laboratory: ['', ],
      Results: ['', ],
      SpecimenType: ['', ]
        });
    }

    
    addLaboratoryOrder(testCode, fastingRequired, Order, Laboratory, Results, SpecimenType): void {
        this.laboratoryOrderService
        .addLaboratoryOrder(testCode, fastingRequired, Order, Laboratory, Results, SpecimenType)
            .subscribe(() => {
                this.router.navigate(['/indexLaboratoryOrder']);
            });
    }

    ngOnInit(): void {
    }
}