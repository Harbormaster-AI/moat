import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { InspectionCharacteristicService } from '../../../services/InspectionCharacteristic.service';
import { InspectionCharacteristic } from '../../../models/InspectionCharacteristic';
import { SubBaseComponent } from '../../InspectionCharacteristic/sub.base.component';

@Component({
    selector: 'app-create-inspectionCharacteristic',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateInspectionCharacteristicComponent extends SubBaseComponent implements OnInit {

    title = 'Add InspectionCharacteristic';

    inspectionCharacteristicForm: FormGroup;
    inspectionCharacteristic: InspectionCharacteristic;

    constructor( http: HttpClient,
        private inspectionCharacteristicService: InspectionCharacteristicService,
        private fb: FormBuilder,
        private router: Router
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

    
    addInspectionCharacteristic(characteristicCode, name, lowerSpecLimit, upperSpecLimit, target, InspectionPlan, MeasurementType): void {
        this.inspectionCharacteristicService
        .addInspectionCharacteristic(characteristicCode, name, lowerSpecLimit, upperSpecLimit, target, InspectionPlan, MeasurementType)
            .subscribe(() => {
                this.router.navigate(['/indexInspectionCharacteristic']);
            });
    }

    ngOnInit(): void {
    }
}