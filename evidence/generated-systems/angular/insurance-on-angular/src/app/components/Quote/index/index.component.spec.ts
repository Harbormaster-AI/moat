
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexQuoteComponent } from './index.component';
import { QuoteService } from '../../../services/Quote.service';

describe('IndexQuoteComponent', () => {
  let component: IndexQuoteComponent;
  let fixture: ComponentFixture<IndexQuoteComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexQuoteComponent
      ],
      providers: [
        QuoteService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexQuoteComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});