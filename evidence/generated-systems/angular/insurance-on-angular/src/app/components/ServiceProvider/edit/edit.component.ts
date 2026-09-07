import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ServiceProviderService } from '../../../services/ServiceProvider.service';
import { SubBaseComponent } from '../../ServiceProvider/sub.base.component';


@Component({
    selector: 'app-edit-serviceProvider',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditServiceProviderComponent extends SubBaseComponent implements OnInit {

    title = 'Edit ServiceProvider';

    serviceProviderForm: FormGroup;
    serviceProvider: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ServiceProviderService,
        private fb: FormBuilder
) {
        super(http);
        this.serviceProviderForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      taxId: ['', Validators.required],
      Claims: ['', ],
      ProviderType: ['', ],
      NetworkStatus: ['', ]
        });
    }

    
    updateServiceProvider(name, taxId, Claims, ProviderType, NetworkStatus): void {
        this.route.params.subscribe((params) => {

                        this.service.updateServiceProvider(name, taxId, Claims, ProviderType, NetworkStatus, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexServiceProvider']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getServiceProvider(params['id']).subscribe(res => {
                this.serviceProvider = res;
            });
        });
    }
}