import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { PositionService } from '../../../services/Position.service';
import { SubBaseComponent } from '../../Position/sub.base.component';


@Component({
    selector: 'app-edit-position',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditPositionComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Position';

    positionForm: FormGroup;
    position: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: PositionService,
        private fb: FormBuilder
) {
        super(http);
        this.positionForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  positionCode: ['', Validators.required],
      fte: ['', Validators.required],
      Department: ['', ],
      JobProfile: ['', ],
      CostCenter: ['', ],
      Location: ['', ],
      ManagerPosition: ['', ],
      DirectReports: ['', ],
      Assignments: ['', ],
      Status: ['', ],
      WorkLocationType: ['', ]
        });
    }

    
    updatePosition(positionCode, fte, Department, JobProfile, CostCenter, Location, ManagerPosition, DirectReports, Assignments, Status, WorkLocationType): void {
        this.route.params.subscribe((params) => {

                        this.service.updatePosition(positionCode, fte, Department, JobProfile, CostCenter, Location, ManagerPosition, DirectReports, Assignments, Status, WorkLocationType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexPosition']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getPosition(params['id']).subscribe(res => {
                this.position = res;
            });
        });
    }
}