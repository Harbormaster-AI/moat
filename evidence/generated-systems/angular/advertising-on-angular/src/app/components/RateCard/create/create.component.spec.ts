
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateRateCardComponent } from './create.component';
import { RateCardService } from '../../../services/RateCard.service';
import { Router } from '@angular/router';

describe('CreateRateCardComponent', () => {
  let component: CreateRateCardComponent;
  let fixture: ComponentFixture<CreateRateCardComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateRateCardComponent
      ],
      providers: [
        RateCardService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateRateCardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});