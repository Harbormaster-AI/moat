import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { OperationService } from '../../../services/Operation.service';
import { Operation } from '../../../models/Operation';
import { SubBaseComponent } from '../../Operation/sub.base.component';

@Component({
    selector: 'app-create-operation',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateOperationComponent extends SubBaseComponent implements OnInit {

    title = 'Add Operation';

    operationForm: FormGroup;
    operation: Operation;

    constructor( http: HttpClient,
        private operationService: OperationService,
        private fb: FormBuilder,
        private router: Router
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

    
    addOperation(operationNumber, name, setupTime, standardCycleTime, Routing, WorkCenter, InspectionPlan, OperationType): void {
        this.operationService
        .addOperation(operationNumber, name, setupTime, standardCycleTime, Routing, WorkCenter, InspectionPlan, OperationType)
            .subscribe(() => {
                this.router.navigate(['/indexOperation']);
            });
    }

    ngOnInit(): void {
    }
}