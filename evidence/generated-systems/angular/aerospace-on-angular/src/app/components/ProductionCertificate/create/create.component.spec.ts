
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateProductionCertificateComponent } from './create.component';
import { ProductionCertificateService } from '../../../services/ProductionCertificate.service';
import { Router } from '@angular/router';

describe('CreateProductionCertificateComponent', () => {
  let component: CreateProductionCertificateComponent;
  let fixture: ComponentFixture<CreateProductionCertificateComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateProductionCertificateComponent
      ],
      providers: [
        ProductionCertificateService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateProductionCertificateComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});