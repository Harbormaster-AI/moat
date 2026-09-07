
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateObligationComponent } from './create.component';
import { ObligationService } from '../../../services/Obligation.service';
import { Router } from '@angular/router';

describe('CreateObligationComponent', () => {
  let component: CreateObligationComponent;
  let fixture: ComponentFixture<CreateObligationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateObligationComponent
      ],
      providers: [
        ObligationService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateObligationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});