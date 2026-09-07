import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AdjusterService } from '../../../services/Adjuster.service';
import { SubBaseComponent } from '../../Adjuster/sub.base.component';


@Component({
    selector: 'app-edit-adjuster',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAdjusterComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Adjuster';

    adjusterForm: FormGroup;
    adjuster: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: AdjusterService,
        private fb: FormBuilder
) {
        super(http);
        this.adjusterForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  firstName: ['', Validators.required],
      lastName: ['', Validators.required],
      licenseNumber: ['', Validators.required],
      Claims: ['', ],
      ServiceProviders: ['', ],
      AdjusterType: ['', ]
        });
    }

    
    updateAdjuster(firstName, lastName, licenseNumber, Claims, ServiceProviders, AdjusterType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAdjuster(firstName, lastName, licenseNumber, Claims, ServiceProviders, AdjusterType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAdjuster']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAdjuster(params['id']).subscribe(res => {
                this.adjuster = res;
            });
        });
    }
}