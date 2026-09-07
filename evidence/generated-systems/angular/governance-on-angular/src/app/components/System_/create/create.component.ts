import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { System_Service } from '../../../services/System_.service';
import { System_ } from '../../../models/System_';
import { SubBaseComponent } from '../../System_/sub.base.component';

@Component({
    selector: 'app-create-system_',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateSystem_Component extends SubBaseComponent implements OnInit {

    title = 'Add System_';

    system_Form: FormGroup;
    system_: System_;

    constructor( http: HttpClient,
        private system_Service: System_Service,
        private fb: FormBuilder,
        private router: Router
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

    
    addSystem_(name, ownerDepartment, ProcessingActivities, RecordsRepositories, SystemType): void {
        this.system_Service
        .addSystem_(name, ownerDepartment, ProcessingActivities, RecordsRepositories, SystemType)
            .subscribe(() => {
                this.router.navigate(['/indexSystem_']);
            });
    }

    ngOnInit(): void {
    }
}