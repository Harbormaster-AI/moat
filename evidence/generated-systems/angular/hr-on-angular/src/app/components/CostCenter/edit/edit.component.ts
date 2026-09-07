import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { CostCenterService } from '../../../services/CostCenter.service';
import { SubBaseComponent } from '../../CostCenter/sub.base.component';


@Component({
    selector: 'app-edit-costCenter',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditCostCenterComponent extends SubBaseComponent implements OnInit {

    title = 'Edit CostCenter';

    costCenterForm: FormGroup;
    costCenter: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: CostCenterService,
        private fb: FormBuilder
) {
        super(http);
        this.costCenterForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  code: ['', Validators.required],
      name: ['', Validators.required],
      Organization: ['', ],
      Departments: ['', ],
      Positions: ['', ],
      Employees: ['', ]
        });
    }

    
    updateCostCenter(code, name, Organization, Departments, Positions, Employees): void {
        this.route.params.subscribe((params) => {

                        this.service.updateCostCenter(code, name, Organization, Departments, Positions, Employees, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexCostCenter']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getCostCenter(params['id']).subscribe(res => {
                this.costCenter = res;
            });
        });
    }
}