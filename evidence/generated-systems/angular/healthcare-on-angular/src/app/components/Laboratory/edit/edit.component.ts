import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { LaboratoryService } from '../../../services/Laboratory.service';
import { SubBaseComponent } from '../../Laboratory/sub.base.component';


@Component({
    selector: 'app-edit-laboratory',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditLaboratoryComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Laboratory';

    laboratoryForm: FormGroup;
    laboratory: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: LaboratoryService,
        private fb: FormBuilder
) {
        super(http);
        this.laboratoryForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      cliaNumber: ['', Validators.required],
      Facility: ['', ],
      LaboratoryOrders: ['', ],
      LabResults: ['', ]
        });
    }

    
    updateLaboratory(name, cliaNumber, Facility, LaboratoryOrders, LabResults): void {
        this.route.params.subscribe((params) => {

                        this.service.updateLaboratory(name, cliaNumber, Facility, LaboratoryOrders, LabResults, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexLaboratory']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getLaboratory(params['id']).subscribe(res => {
                this.laboratory = res;
            });
        });
    }
}