import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { System_Service } from '../../../services/System_.service';
import { SubBaseComponent } from '../../System_/sub.base.component';


@Component({
    selector: 'app-edit-system_',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditSystem_Component extends SubBaseComponent implements OnInit {

    title = 'Edit System_';

    system_Form: FormGroup;
    system_: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: System_Service,
        private fb: FormBuilder
) {
        super(http);
        this.system_Form = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      ownerDepartment: ['', Validators.required],
      ProcessingActivities: ['', ],
      RecordsRepositories: ['', ],
      SystemType: ['', ]
        });
    }

    
    updateSystem_(name, ownerDepartment, ProcessingActivities, RecordsRepositories, SystemType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateSystem_(name, ownerDepartment, ProcessingActivities, RecordsRepositories, SystemType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexSystem_']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getSystem_(params['id']).subscribe(res => {
                this.system_ = res;
            });
        });
    }
}