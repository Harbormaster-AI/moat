import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { OperationService } from '../../../services/Operation.service';
import { SubBaseComponent } from '../../Operation/sub.base.component';


@Component({
    selector: 'app-edit-operation',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditOperationComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Operation';

    operationForm: FormGroup;
    operation: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: OperationService,
        private fb: FormBuilder
) {
        super(http);
        this.operationForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  operationNumber: ['', Validators.required],
      name: ['', Validators.required],
      setupTime: ['', Validators.required],
      standardCycleTime: ['', Validators.required],
      Routing: ['', ],
      WorkCenter: ['', ],
      InspectionPlan: ['', ],
      OperationType: ['', ]
        });
    }

    
    updateOperation(operationNumber, name, setupTime, standardCycleTime, Routing, WorkCenter, InspectionPlan, OperationType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateOperation(operationNumber, name, setupTime, standardCycleTime, Routing, WorkCenter, InspectionPlan, OperationType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexOperation']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getOperation(params['id']).subscribe(res => {
                this.operation = res;
            });
        });
    }
}