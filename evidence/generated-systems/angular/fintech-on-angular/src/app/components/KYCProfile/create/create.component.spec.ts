
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateKYCProfileComponent } from './create.component';
import { KYCProfileService } from '../../../services/KYCProfile.service';
import { Router } from '@angular/router';

describe('CreateKYCProfileComponent', () => {
  let component: CreateKYCProfileComponent;
  let fixture: ComponentFixture<CreateKYCProfileComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateKYCProfileComponent
      ],
      providers: [
        KYCProfileService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateKYCProfileComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});