
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateAttestationComponent } from './create.component';
import { AttestationService } from '../../../services/Attestation.service';
import { Router } from '@angular/router';

describe('CreateAttestationComponent', () => {
  let component: CreateAttestationComponent;
  let fixture: ComponentFixture<CreateAttestationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateAttestationComponent
      ],
      providers: [
        AttestationService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateAttestationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});