import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { CabinLayoutService } from '../../../services/CabinLayout.service';
import { SubBaseComponent } from '../../CabinLayout/sub.base.component';


@Component({
    selector: 'app-edit-cabinLayout',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditCabinLayoutComponent extends SubBaseComponent implements OnInit {

    title = 'Edit CabinLayout';

    cabinLayoutForm: FormGroup;
    cabinLayout: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: CabinLayoutService,
        private fb: FormBuilder
) {
        super(http);
        this.cabinLayoutForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  layoutCode: ['', Validators.required],
      totalSeats: ['', Validators.required],
      classLayout: ['', Validators.required],
      Variant: ['', ],
      Aircraft: ['', ],
      Options: ['', ]
        });
    }

    
    updateCabinLayout(layoutCode, totalSeats, classLayout, Variant, Aircraft, Options): void {
        this.route.params.subscribe((params) => {

                        this.service.updateCabinLayout(layoutCode, totalSeats, classLayout, Variant, Aircraft, Options, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexCabinLayout']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getCabinLayout(params['id']).subscribe(res => {
                this.cabinLayout = res;
            });
        });
    }
}