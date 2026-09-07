import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { DeviceCriterionService } from '../../../services/DeviceCriterion.service';
import { DeviceCriterion } from '../../../models/DeviceCriterion';
import { SubBaseComponent } from '../../DeviceCriterion/sub.base.component';

@Component({
    selector: 'app-create-deviceCriterion',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateDeviceCriterionComponent extends SubBaseComponent implements OnInit {

    title = 'Add DeviceCriterion';

    deviceCriterionForm: FormGroup;
    deviceCriterion: DeviceCriterion;

    constructor( http: HttpClient,
        private deviceCriterionService: DeviceCriterionService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.deviceCriterionForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  TargetingProfile: ['', ],
      DeviceType: ['', ],
      PlatformType: ['', ],
      Operator: ['', ]
        });
    }

    
    addDeviceCriterion(TargetingProfile, DeviceType, PlatformType, Operator): void {
        this.deviceCriterionService
        .addDeviceCriterion(TargetingProfile, DeviceType, PlatformType, Operator)
            .subscribe(() => {
                this.router.navigate(['/indexDeviceCriterion']);
            });
    }

    ngOnInit(): void {
    }
}