
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexDataProviderComponent } from './index.component';
import { DataProviderService } from '../../../services/DataProvider.service';

describe('IndexDataProviderComponent', () => {
  let component: IndexDataProviderComponent;
  let fixture: ComponentFixture<IndexDataProviderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexDataProviderComponent
      ],
      providers: [
        DataProviderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexDataProviderComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});