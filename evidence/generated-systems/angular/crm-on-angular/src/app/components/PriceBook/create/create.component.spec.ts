
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreatePriceBookComponent } from './create.component';
import { PriceBookService } from '../../../services/PriceBook.service';
import { Router } from '@angular/router';

describe('CreatePriceBookComponent', () => {
  let component: CreatePriceBookComponent;
  let fixture: ComponentFixture<CreatePriceBookComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreatePriceBookComponent
      ],
      providers: [
        PriceBookService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreatePriceBookComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});