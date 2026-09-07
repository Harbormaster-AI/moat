import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { SalaryComponentService } from '../../../services/SalaryComponent.service';
import { SubBaseComponent } from '../../SalaryComponent/sub.base.component';


@Component({
    selector: 'app-edit-salaryComponent',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditSalaryComponentComponent extends SubBaseComponent implements OnInit {

    title = 'Edit SalaryComponent';

    salaryComponentForm: FormGroup;
    salaryComponent: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: SalaryComponentService,
        private fb: FormBuilder
) {
        super(http);
        this.salaryComponentForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  amount: ['', Validators.required],
      recurring: ['', Validators.required],
      CompensationPackage: ['', ],
      ComponentType: ['', ]
        });
    }

    
    updateSalaryComponent(amount, recurring, CompensationPackage, ComponentType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateSalaryComponent(amount, recurring, CompensationPackage, ComponentType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexSalaryComponent']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getSalaryComponent(params['id']).subscribe(res => {
                this.salaryComponent = res;
            });
        });
    }
}