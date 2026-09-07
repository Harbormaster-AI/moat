import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { DischargeService } from '../../../services/Discharge.service';
import { SubBaseComponent } from '../../Discharge/sub.base.component';


@Component({
    selector: 'app-edit-discharge',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditDischargeComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Discharge';

    dischargeForm: FormGroup;
    discharge: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: DischargeService,
        private fb: FormBuilder
) {
        super(http);
        this.dischargeForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  dischargeDateTime: ['', Validators.required],
      Encounter: ['', ],
      Disposition: ['', ]
        });
    }

    
    updateDischarge(dischargeDateTime, Encounter, Disposition): void {
        this.route.params.subscribe((params) => {

                        this.service.updateDischarge(dischargeDateTime, Encounter, Disposition, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexDischarge']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getDischarge(params['id']).subscribe(res => {
                this.discharge = res;
            });
        });
    }
}