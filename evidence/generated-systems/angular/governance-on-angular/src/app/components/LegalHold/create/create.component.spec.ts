
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateLegalHoldComponent } from './create.component';
import { LegalHoldService } from '../../../services/LegalHold.service';
import { Router } from '@angular/router';

describe('CreateLegalHoldComponent', () => {
  let component: CreateLegalHoldComponent;
  let fixture: ComponentFixture<CreateLegalHoldComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateLegalHoldComponent
      ],
      providers: [
        LegalHoldService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateLegalHoldComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});