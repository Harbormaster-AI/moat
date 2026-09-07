import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { InspectionCharacteristicService } from '../../../services/InspectionCharacteristic.service';
import { SubBaseComponent } from '../../InspectionCharacteristic/sub.base.component';


@Component({
    selector: 'app-edit-inspectionCharacteristic',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditInspectionCharacteristicComponent extends SubBaseComponent implements OnInit {

    title = 'Edit InspectionCharacteristic';

    inspectionCharacteristicForm: FormGroup;
    inspectionCharacteristic: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: InspectionCharacteristicService,
        private fb: FormBuilder
) {
        super(http);
        this.inspectionCharacteristicForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  characteristicCode: ['', Validators.required],
      name: ['', Validators.required],
      lowerSpecLimit: ['', Validators.required],
      upperSpecLimit: ['', Validators.required],
      target: ['', Validators.required],
      InspectionPlan: ['', ],
      MeasurementType: ['', ]
        });
    }

    
    updateInspectionCharacteristic(characteristicCode, name, lowerSpecLimit, upperSpecLimit, target, InspectionPlan, MeasurementType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateInspectionCharacteristic(characteristicCode, name, lowerSpecLimit, upperSpecLimit, target, InspectionPlan, MeasurementType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexInspectionCharacteristic']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getInspectionCharacteristic(params['id']).subscribe(res => {
                this.inspectionCharacteristic = res;
            });
        });
    }
}