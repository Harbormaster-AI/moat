import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ServiceProviderService } from '../../../services/ServiceProvider.service';
import { ServiceProvider } from '../../../models/ServiceProvider';
import { SubBaseComponent } from '../../ServiceProvider/sub.base.component';

@Component({
    selector: 'app-create-serviceProvider',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateServiceProviderComponent extends SubBaseComponent implements OnInit {

    title = 'Add ServiceProvider';

    serviceProviderForm: FormGroup;
    serviceProvider: ServiceProvider;

    constructor( http: HttpClient,
        private serviceProviderService: ServiceProviderService,
        private fb: FormBuilder,
        private router: Router
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

    
    addServiceProvider(name, taxId, Claims, ProviderType, NetworkStatus): void {
        this.serviceProviderService
        .addServiceProvider(name, taxId, Claims, ProviderType, NetworkStatus)
            .subscribe(() => {
                this.router.navigate(['/indexServiceProvider']);
            });
    }

    ngOnInit(): void {
    }
}