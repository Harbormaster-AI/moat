
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexFXQuoteComponent } from './index.component';
import { FXQuoteService } from '../../../services/FXQuote.service';

describe('IndexFXQuoteComponent', () => {
  let component: IndexFXQuoteComponent;
  let fixture: ComponentFixture<IndexFXQuoteComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexFXQuoteComponent
      ],
      providers: [
        FXQuoteService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexFXQuoteComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});