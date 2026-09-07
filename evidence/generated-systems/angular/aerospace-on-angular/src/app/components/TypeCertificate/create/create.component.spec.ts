
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateTypeCertificateComponent } from './create.component';
import { TypeCertificateService } from '../../../services/TypeCertificate.service';
import { Router } from '@angular/router';

describe('CreateTypeCertificateComponent', () => {
  let component: CreateTypeCertificateComponent;
  let fixture: ComponentFixture<CreateTypeCertificateComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateTypeCertificateComponent
      ],
      providers: [
        TypeCertificateService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateTypeCertificateComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});