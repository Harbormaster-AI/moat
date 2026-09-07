import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { DistributorService } from '../../../services/Distributor.service';
import { SubBaseComponent } from '../../Distributor/sub.base.component';


@Component({
    selector: 'app-edit-distributor',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditDistributorComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Distributor';

    distributorForm: FormGroup;
    distributor: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: DistributorService,
        private fb: FormBuilder
) {
        super(http);
        this.distributorForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      licenseNumber: ['', Validators.required],
      region: ['', Validators.required],
      Insurers: ['', ],
      Agents: ['', ],
      Policies: ['', ],
      DistributorType: ['', ]
        });
    }

    
    updateDistributor(name, licenseNumber, region, Insurers, Agents, Policies, DistributorType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateDistributor(name, licenseNumber, region, Insurers, Agents, Policies, DistributorType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexDistributor']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getDistributor(params['id']).subscribe(res => {
                this.distributor = res;
            });
        });
    }
}