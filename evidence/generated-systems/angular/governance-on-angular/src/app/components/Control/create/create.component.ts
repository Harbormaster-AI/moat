import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ControlService } from '../../../services/Control.service';
import { Control } from '../../../models/Control';
import { SubBaseComponent } from '../../Control/sub.base.component';

@Component({
    selector: 'app-create-control',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateControlComponent extends SubBaseComponent implements OnInit {

    title = 'Add Control';

    controlForm: FormGroup;
    control: Control;

    constructor( http: HttpClient,
        private controlService: ControlService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.controlForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      objective: ['', Validators.required],
      ownerDepartment: ['', Validators.required],
      Policy: ['', ],
      ControlTests: ['', ],
      Evidence: ['', ],
      Risks: ['', ],
      Obligations: ['', ],
      Procedures: ['', ],
      Issues: ['', ],
      ControlType: ['', ],
      Frequency: ['', ],
      Status: ['', ]
        });
    }

    
    addControl(name, objective, ownerDepartment, Policy, ControlTests, Evidence, Risks, Obligations, Procedures, Issues, ControlType, Frequency, Status): void {
        this.controlService
        .addControl(name, objective, ownerDepartment, Policy, ControlTests, Evidence, Risks, Obligations, Procedures, Issues, ControlType, Frequency, Status)
            .subscribe(() => {
                this.router.navigate(['/indexControl']);
            });
    }

    ngOnInit(): void {
    }
}