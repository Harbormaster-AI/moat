
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateFXQuoteComponent } from './create.component';
import { FXQuoteService } from '../../../services/FXQuote.service';
import { Router } from '@angular/router';

describe('CreateFXQuoteComponent', () => {
  let component: CreateFXQuoteComponent;
  let fixture: ComponentFixture<CreateFXQuoteComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateFXQuoteComponent
      ],
      providers: [
        FXQuoteService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateFXQuoteComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});