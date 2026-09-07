
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateFXDealComponent } from './create.component';
import { FXDealService } from '../../../services/FXDeal.service';
import { Router } from '@angular/router';

describe('CreateFXDealComponent', () => {
  let component: CreateFXDealComponent;
  let fixture: ComponentFixture<CreateFXDealComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateFXDealComponent
      ],
      providers: [
        FXDealService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateFXDealComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});