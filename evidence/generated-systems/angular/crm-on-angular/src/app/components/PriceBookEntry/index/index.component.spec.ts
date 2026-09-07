
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexPriceBookEntryComponent } from './index.component';
import { PriceBookEntryService } from '../../../services/PriceBookEntry.service';

describe('IndexPriceBookEntryComponent', () => {
  let component: IndexPriceBookEntryComponent;
  let fixture: ComponentFixture<IndexPriceBookEntryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexPriceBookEntryComponent
      ],
      providers: [
        PriceBookEntryService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexPriceBookEntryComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});