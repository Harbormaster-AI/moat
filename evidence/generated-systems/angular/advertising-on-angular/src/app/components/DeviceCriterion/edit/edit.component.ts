import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { DeviceCriterionService } from '../../../services/DeviceCriterion.service';
import { SubBaseComponent } from '../../DeviceCriterion/sub.base.component';


@Component({
    selector: 'app-edit-deviceCriterion',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditDeviceCriterionComponent extends SubBaseComponent implements OnInit {

    title = 'Edit DeviceCriterion';

    deviceCriterionForm: FormGroup;
    deviceCriterion: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: DeviceCriterionService,
        private fb: FormBuilder
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

    
    updateDeviceCriterion(TargetingProfile, DeviceType, PlatformType, Operator): void {
        this.route.params.subscribe((params) => {

                        this.service.updateDeviceCriterion(TargetingProfile, DeviceType, PlatformType, Operator, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexDeviceCriterion']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getDeviceCriterion(params['id']).subscribe(res => {
                this.deviceCriterion = res;
            });
        });
    }
}