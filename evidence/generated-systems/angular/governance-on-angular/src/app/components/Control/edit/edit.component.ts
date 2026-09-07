import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ControlService } from '../../../services/Control.service';
import { SubBaseComponent } from '../../Control/sub.base.component';


@Component({
    selector: 'app-edit-control',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditControlComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Control';

    controlForm: FormGroup;
    control: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ControlService,
        private fb: FormBuilder
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

    
    updateControl(name, objective, ownerDepartment, Policy, ControlTests, Evidence, Risks, Obligations, Procedures, Issues, ControlType, Frequency, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateControl(name, objective, ownerDepartment, Policy, ControlTests, Evidence, Risks, Obligations, Procedures, Issues, ControlType, Frequency, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexControl']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getControl(params['id']).subscribe(res => {
                this.control = res;
            });
        });
    }
}