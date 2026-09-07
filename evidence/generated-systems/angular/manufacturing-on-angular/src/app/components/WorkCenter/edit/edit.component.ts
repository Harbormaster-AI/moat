import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { WorkCenterService } from '../../../services/WorkCenter.service';
import { SubBaseComponent } from '../../WorkCenter/sub.base.component';


@Component({
    selector: 'app-edit-workCenter',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditWorkCenterComponent extends SubBaseComponent implements OnInit {

    title = 'Edit WorkCenter';

    workCenterForm: FormGroup;
    workCenter: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: WorkCenterService,
        private fb: FormBuilder
) {
        super(http);
        this.workCenterForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      code: ['', Validators.required],
      capacityPerHour: ['', Validators.required],
      oeeTarget: ['', Validators.required],
      ProductionLine: ['', ],
      Assets: ['', ],
      MaintenanceOrders: ['', ],
      WorkCenterType: ['', ]
        });
    }

    
    updateWorkCenter(name, code, capacityPerHour, oeeTarget, ProductionLine, Assets, MaintenanceOrders, WorkCenterType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateWorkCenter(name, code, capacityPerHour, oeeTarget, ProductionLine, Assets, MaintenanceOrders, WorkCenterType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexWorkCenter']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getWorkCenter(params['id']).subscribe(res => {
                this.workCenter = res;
            });
        });
    }
}