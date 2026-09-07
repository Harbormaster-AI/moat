import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { LaboratoryOrderService } from '../../../services/LaboratoryOrder.service';
import { SubBaseComponent } from '../../LaboratoryOrder/sub.base.component';


@Component({
    selector: 'app-edit-laboratoryOrder',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditLaboratoryOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Edit LaboratoryOrder';

    laboratoryOrderForm: FormGroup;
    laboratoryOrder: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: LaboratoryOrderService,
        private fb: FormBuilder
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

    
    updateLaboratoryOrder(testCode, fastingRequired, Order, Laboratory, Results, SpecimenType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateLaboratoryOrder(testCode, fastingRequired, Order, Laboratory, Results, SpecimenType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexLaboratoryOrder']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getLaboratoryOrder(params['id']).subscribe(res => {
                this.laboratoryOrder = res;
            });
        });
    }
}