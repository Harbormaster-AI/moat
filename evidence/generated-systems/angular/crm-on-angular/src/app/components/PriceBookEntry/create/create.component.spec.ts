
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreatePriceBookEntryComponent } from './create.component';
import { PriceBookEntryService } from '../../../services/PriceBookEntry.service';
import { Router } from '@angular/router';

describe('CreatePriceBookEntryComponent', () => {
  let component: CreatePriceBookEntryComponent;
  let fixture: ComponentFixture<CreatePriceBookEntryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreatePriceBookEntryComponent
      ],
      providers: [
        PriceBookEntryService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreatePriceBookEntryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});